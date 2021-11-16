package mesh

//import (
//	"context"
//	"fmt"
//	"time"
//
//	"github.com/enjoypi/god/pb"
//	sc "github.com/enjoypi/gostatechart"
//	etcdclient "go.etcd.io/etcd/clientv3"
//)
//
//func (m *main) onServiceInfo(event sc.Event) sc.Event {
//	info := event.(*pb.ServiceInfo)
//	key := fmt.Sprintf("%s/%d", m.Mesh.Path, info.ServiceType)
//
//	// waiting for reconnecting
//	var client *etcdclient.Client
//	for client = m.Client; client == nil; client = m.Client {
//		time.Sleep(100 * time.Millisecond)
//	}
//	_, err := client.Put(context.Background(), key, m.Mesh.AdvertiseAddress)
//	return err
//	//if err != nil {
//	//
//	//}
//}
