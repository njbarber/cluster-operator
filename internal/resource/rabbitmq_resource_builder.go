// RabbitMQ Cluster Operator
//
// Copyright 2020 VMware, Inc. All Rights Reserved.
//
// This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.
//
// This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
//

package resource

import (
	rabbitmqv1beta2 "github.com/rabbitmq/cluster-operator/api/v1beta2"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RabbitmqResourceBuilder struct {
	Instance *rabbitmqv1beta2.RabbitmqCluster
	Scheme   *runtime.Scheme
}

type ResourceBuilder interface {
	Build() (client.Object, error)
	Update(client.Object) error
	UpdateMayRequireStsRecreate() bool
}

func (builder *RabbitmqResourceBuilder) ResourceBuilders() ([]ResourceBuilder, error) {
	return []ResourceBuilder{
		builder.HeadlessService(),
		builder.Service(),
		builder.ErlangCookie(),
		builder.DefaultUserSecret(),
		builder.RabbitmqPluginsConfigMap(),
		builder.ServerConfigMap(),
		builder.ServiceAccount(),
		builder.Role(),
		builder.RoleBinding(),
		builder.StatefulSet(),
	}, nil
}
