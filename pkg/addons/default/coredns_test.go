package defaultaddons_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/weaveworks/eksctl/pkg/addons/default"

	"github.com/weaveworks/eksctl/pkg/testutils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ = Describe("default addons - coredns", func() {
	Describe("can update coredns add-on", func() {
		var (
			rawClient *testutils.FakeRawClient
			ct        *testutils.CollectionTracker
		)

		It("can load sample for 1.10 and create objests that don't exist", func() {
			sampleAddons := testutils.LoadSamples("testdata/sample-1.10.json")

			rawClient = testutils.NewFakeRawClient()

			rawClient.AssumeObjectsMissing = true

			for _, item := range sampleAddons {
				rc, err := rawClient.NewRawResource(runtime.RawExtension{Object: item})
				Expect(err).ToNot(HaveOccurred())
				_, err = rc.CreateOrReplace()
				Expect(err).ToNot(HaveOccurred())
			}

			ct = rawClient.Collection

			Expect(ct.Updated()).To(BeEmpty())
			Expect(ct.Created()).ToNot(BeEmpty())
			Expect(ct.CreatedItems()).To(HaveLen(6))

			Expect(ct.CreatedItems()).To(HaveLen(6))
		})

		It("has newly created objects", func() {
			rawClient.ClientSetUseUpdatedObjects = false // must be set for initial verification, and for subsequent UpdateAWSNode

			kubeDNS, err := rawClient.ClientSet().AppsV1().Deployments(metav1.NamespaceSystem).Get(KubeDNS, metav1.GetOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(kubeDNS.Spec.Template.Spec.Containers).To(HaveLen(3))
			Expect(kubeDNS.Spec.Template.Spec.Containers[0].Image).To(
				Equal("602401143452.dkr.ecr.eu-west-2.amazonaws.com/eks/kube-dns/kube-dns:1.14.10"),
			)
			Expect(kubeDNS.Spec.Template.Spec.Containers[1].Image).To(
				Equal("602401143452.dkr.ecr.eu-west-2.amazonaws.com/eks/kube-dns/dnsmasq-nanny:1.14.10"),
			)
			Expect(kubeDNS.Spec.Template.Spec.Containers[2].Image).To(
				Equal("602401143452.dkr.ecr.eu-west-2.amazonaws.com/eks/kube-dns/sidecar:1.14.10"),
			)
		})

		It("can update 1.10 sample to latest", func() {
			rawClient.AssumeObjectsMissing = false

			err := InstallCoreDNS(rawClient, "eu-west-2", 1*time.Minute)
			Expect(err).ToNot(HaveOccurred())
			Expect(rawClient.Collection.UpdatedItems()).To(HaveLen(4))
			Expect(rawClient.Collection.CreatedItems()).To(HaveLen(6))

			rawClient.ClientSetUseUpdatedObjects = true // for verification of updated objects

			awsNode, err := rawClient.ClientSet().AppsV1().DaemonSets(metav1.NamespaceSystem).Get(AWSNode, metav1.GetOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(awsNode.Spec.Template.Spec.Containers).To(HaveLen(1))
			Expect(awsNode.Spec.Template.Spec.Containers[0].Image).To(
				Equal("602401143452.dkr.ecr.eu-west-2.amazonaws.com/amazon-k8s-cni:v1.3.2"),
			)

			rawClient.ClearUpdated()
		})

		// It("can update 1.10 sample for different region", func() {
		// 	rawClient.ClientSetUseUpdatedObjects = false // must be set for subsequent UpdateAWSNode

		// 	err := UpdateAWSNode(rawClient, "us-east-1")
		// 	Expect(err).ToNot(HaveOccurred())

		// 	rawClient.ClientSetUseUpdatedObjects = true // for verification of updated objects

		// 	awsNode, err := rawClient.ClientSet().AppsV1().DaemonSets(metav1.NamespaceSystem).Get(AWSNode, metav1.GetOptions{})
		// 	Expect(err).ToNot(HaveOccurred())
		// 	Expect(awsNode.Spec.Template.Spec.Containers).To(HaveLen(1))
		// 	Expect(awsNode.Spec.Template.Spec.Containers[0].Image).To(
		// 		Equal("602401143452.dkr.ecr.us-east-1.amazonaws.com/amazon-k8s-cni:v1.3.2"),
		// 	)
		// })

	})
})
