package client

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func restClientDemo() {
	// 从~/.kube/config读取配置并反序列化
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic("can't create client, err:" + err.Error())
	}
	// 设置api版本
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// 根据配置生成client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic("can't create client, err:" + err.Error())
	}
	pod := &corev1.Pod{}
	// 基于client获取资源，标识一个资源的请求信息，需要 namespace/kind/name，
	//比如说default命名空间下 name为nginx-pod的pod
	err = restClient.Get().Namespace("default").Resource("pods").
		Name("nginx-pod").Do(context.TODO()).Into(pod)

	if err != nil {
		panic("can't get pod, err:" + err.Error())
	} else {
		fmt.Printf("%s/%s\n", pod.Namespace, pod.Name)
	}
}

func restClientGetPodListDemo() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic("can't create client, err:" + err.Error())
	}
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// 根据配置生成client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic("can't create client, err:" + err.Error())
	}

	// 获取pod 列表
	podList := &corev1.PodList{}
	err = restClient.Get().Namespace(metav1.NamespaceAll).Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 10}, scheme.ParameterCodec).
		Do(context.TODO()).Into(podList)
	if err != nil {
		return
	}
	for _, pod := range podList.Items {
		fmt.Printf("%s/%s\n", pod.Namespace, pod.Name)
	}
}
