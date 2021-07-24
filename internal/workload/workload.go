package workload

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"git.sr.ht/~spc/go-log"
	api2 "github.com/jakub-dzon/k4e-device-worker/internal/workload/api"
	"github.com/jakub-dzon/k4e-device-worker/internal/workload/network"
	podman2 "github.com/jakub-dzon/k4e-device-worker/internal/workload/podman"
	"github.com/jakub-dzon/k4e-operator/models"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

const nfTableName string = "edge"

type WorkloadManager struct {
	manifestsDir string
	workloads    *workloadWrapper
}

// workloadWrapper manages the workload and its configuration on the device
type workloadWrapper struct {
	workloads *podman2.Podman
	netfilter *network.Netfilter
}

func newWorkloadWrapper() (*workloadWrapper, error) {
	newPodman, err := podman2.NewPodman()
	if err != nil {
		return nil, err
	}
	netfilter, err := network.NewNetfilter()
	if err != nil {
		return nil, err
	}
	return &workloadWrapper{
		workloads: newPodman,
		netfilter: netfilter,
	}, nil
}

func NewWorkloadManager(configDir string) (*WorkloadManager, error) {
	manifestsDir := path.Join(configDir, "manifests")
	if err := os.MkdirAll(manifestsDir, 0755); err != nil {
		return nil, fmt.Errorf("cannot create directory: %w", err)
	}
	wrapper, err := newWorkloadWrapper()
	if err != nil {
		return nil, err
	}
	manager := WorkloadManager{
		manifestsDir: manifestsDir,
		workloads:    wrapper,
	}
	if err := manager.workloads.Init(); err != nil {
		return nil, err
	}
	go func() {
		for {
			err := manager.ensureWorkloadsFromManifestsAreRunning()
			if err != nil {
				log.Error(err)
			}
			time.Sleep(time.Second * 15)
		}
	}()

	return &manager, nil
}

func (w *WorkloadManager) ListWorkloads() ([]api2.WorkloadInfo, error) {
	return w.workloads.List()
}

func (w *WorkloadManager) Update(configuration models.DeviceConfigurationMessage) error {
	workloads := configuration.Workloads
	if len(workloads) == 0 {
		log.Trace("No workloads")

		// Purge all the workloads
		err := w.purgeWorkloads()
		if err != nil {
			return err
		}
		// Remove manifests
		err = w.removeManifests()
		if err != nil {
			return err
		}
		return nil
	}

	for _, workload := range workloads {
		log.Tracef("Deploying workload: %s", workload.Name)
		// TODO: change error handling from fail fast to best effort (deploy as many workloads as possible)
		pod, err := w.toPod(workload)
		if err != nil {
			return err
		}
		manifestPath, err := w.storeManifest(pod)
		if err != nil {
			return err
		}

		err = w.workloads.Remove(workload.Name)
		if err != nil {
			log.Errorf("Error removing workload: %v", err)
			return err
		}
		err = w.workloads.Run(pod, manifestPath)
		if err != nil {
			log.Errorf("Cannot run workload: %v", err)
			return err
		}
	}
	return nil
}

func (w *WorkloadManager) purgeWorkloads() error {
	podList, err := w.workloads.List()
	if err != nil {
		log.Errorf("Cannot list workloads: %v", err)
		return err
	}
	for _, podReport := range podList {
		err := w.workloads.Remove(podReport.Name)
		if err != nil {
			log.Errorf("Error removing workload: %v", err)
			return err
		}
	}
	return nil
}

func (w *WorkloadManager) removeManifests() error {
	manifestInfo, err := ioutil.ReadDir(w.manifestsDir)
	if err != nil {
		return err
	}
	for _, fi := range manifestInfo {
		filePath := path.Join(w.manifestsDir, fi.Name())
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *WorkloadManager) storeManifest(pod *v1.Pod) (string, error) {
	podYaml, err := yaml.Marshal(pod)
	if err != nil {
		return "", err
	}
	fileName := strings.ReplaceAll(pod.Name, " ", "-") + ".yaml"
	filePath := path.Join(w.manifestsDir, fileName)
	err = ioutil.WriteFile(filePath, podYaml, 0640)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (w *WorkloadManager) ensureWorkloadsFromManifestsAreRunning() error {
	manifestInfo, err := ioutil.ReadDir(w.manifestsDir)
	if err != nil {
		return err
	}
	workloads, err := w.workloads.List()
	if err != nil {
		return err
	}
	nameToWorkload := make(map[string]api2.WorkloadInfo)
	for _, workload := range workloads {
		nameToWorkload[workload.Name] = workload
	}
	for _, fi := range manifestInfo {
		filePath := path.Join(w.manifestsDir, fi.Name())
		manifest, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Error(err)
			continue
		}
		pod := &v1.Pod{}
		err = yaml.Unmarshal(manifest, pod)
		if err != nil {
			log.Error(err)
			continue
		}
		if workload, ok := nameToWorkload[pod.Name]; ok {
			if workload.Status != "Running" {
				// Workload is not running - start
				err = w.workloads.Start(pod)
				if err != nil {
					log.Errorf("failed to start workload %s: %v", pod.Name, err)
				}
			}
			continue
		}
		// Workload is not present - run
		err = w.workloads.Run(pod, filePath)
		if err != nil {
			log.Errorf("failed to run workload %s (manifest: %s): %v", pod.Name, filePath, err)
			continue
		}
	}
	return nil
}

func (w *WorkloadManager) toPod(workload *models.Workload) (*v1.Pod, error) {
	podSpec := v1.PodSpec{}
	err := yaml.Unmarshal([]byte(workload.Specification), &podSpec)
	if err != nil {
		return nil, err
	}
	pod := v1.Pod{
		Spec: podSpec,
	}
	pod.Kind = "Pod"
	pod.Name = workload.Name
	return &pod, nil
}

func (ww workloadWrapper) Init() error {
	return ww.netfilter.AddTable(nfTableName)
}

func (ww workloadWrapper) List() ([]api2.WorkloadInfo, error) {
	return ww.workloads.List()
}

func (ww workloadWrapper) Remove(workloadName string) error {
	if err := ww.workloads.Remove(workloadName); err != nil {
		return err
	}
	if err := ww.netfilter.DeleteChain(nfTableName, workloadName); err != nil {
		log.Errorf("failed to delete chain %[1]s from %s table for workload %[1]s", workloadName, nfTableName)
	}
	return nil
}

func (ww workloadWrapper) Run(workload *v1.Pod, manifestPath string) error {
	if err := ww.applyNetworkConfiguration(workload); err != nil {
		return err
	}
	if err := ww.workloads.Run(manifestPath); err != nil {
		return err
	}
	return nil
}

func (ww workloadWrapper) applyNetworkConfiguration(workload *v1.Pod) error {
	hostPorts, err := getHostPorts(workload)
	if err != nil {
		log.Error(err)
		return err
	}
	if len(hostPorts) == 0 {
		return nil
	}
	// skip existence check since chain is not changed if already exists
	if err := ww.netfilter.AddChain(nfTableName, workload.Name); err != nil {
		return fmt.Errorf("failed to create chain for workload %s: %v", workload.Name, err)
	}

	// for workloads, a port will be opened for the pod based on hostPort
	for _, p := range hostPorts {
		rule := fmt.Sprintf("tcp dport %d ct state new,established counter accept", p)
		if err := ww.netfilter.AddRule(nfTableName, workload.Name, rule); err != nil {
			return fmt.Errorf("failed to add rule %s for workload %s: %v", rule, workload.Name, err)
		}
	}
	return nil
}

func (ww workloadWrapper) Start(workload *v1.Pod) error {
	ww.netfilter.DeleteChain(nfTableName, workload.Name)
	if err := ww.applyNetworkConfiguration(workload); err != nil {
		return err
	}
	if err := ww.workloads.Start(workload.Name); err != nil {
		return err
	}
	return nil
}

func getHostPorts(workload *v1.Pod) ([]int32, error) {
	hostPorts := []int32{}
	for _, c := range workload.Spec.Containers {
		for _, p := range c.Ports {
			if p.HostPort > 0 && p.HostPort < 65536 {
				hostPorts = append(hostPorts, p.HostPort)
			} else {
				return nil, fmt.Errorf("illegal host port number %d for container %s in workload %s", p.HostPort, c.Name, workload.Name)
			}
		}
	}
	return hostPorts, nil
}
