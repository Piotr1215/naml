//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package codify

import (
	"bytes"
	"fmt"
	"github.com/kris-nova/logger"
	rbacv1 "k8s.io/api/rbac/v1"
	"text/template"
	"time"
)

type ClusterRole struct {
	i *rbacv1.ClusterRole
}

func NewClusterRole(obj *rbacv1.ClusterRole) *ClusterRole {
	obj.ObjectMeta = cleanObjectMeta(obj.ObjectMeta)
	return &ClusterRole{
		i: obj,
	}
}

func (k ClusterRole) Install() string {
	l := Literal(k.i)
	install := fmt.Sprintf(`
	{{ .Name }}ClusterRole := %s

	_, err = client.RbacV1().ClusterRoles("{{ .Namespace }}").Create(context.TODO(), {{ .Name }}ClusterRole, v1.CreateOptions{})
	if err != nil {
		return err
	}
`, l)

	tpl := template.New(fmt.Sprintf("%s", time.Now().String()))
	tpl.Parse(install)
	buf := &bytes.Buffer{}
	k.i.Name = sanitizeK8sObjectName(k.i.Name)
	err := tpl.Execute(buf, k.i)
	if err != nil {
		logger.Debug(err.Error())
	}
	return alias(buf.String(), "rbacv1")
}

func (k ClusterRole) Uninstall() string {
	uninstall := `
	err = client.RbacV1().ClusterRoles("{{ .Namespace }}").Delete(context.TODO(), "{{ .Name }}", metav1.DeleteOptions{})
	if err != nil {
		return err
	}
 `
	tpl := template.New(fmt.Sprintf("%s", time.Now().String()))
	tpl.Parse(uninstall)
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, k.i)
	if err != nil {
		logger.Debug(err.Error())
	}
	return buf.String()
}