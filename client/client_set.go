package client

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func clientSetDemo() {
	// 读取kubeconfig配置
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic("can't create client, err:" + err.Error())
	}
	// 根据配置生成clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("can't create client, err:" + err.Error())
		return
	}
	// 使用client获取default命名空间下的pod列表
	podList, err := clientSet.CoreV1().Pods(metav1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic("can't get pod, err:" + err.Error())
	}

	// 打印default命名空间下的pod列表
	for _, item := range podList.Items {
		nn, err := cache.MetaNamespaceKeyFunc(&item)
		fmt.Printf("%s, %v\n", nn, err)
	}
}
