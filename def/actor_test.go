package def

//func TestCodec(t *testing.T) {
//	nodeID := NodeID(1)
//	actorID := ActorID(1)
//	fullID := FullID(0x00010001)
//
//	require.Equal(t, fullID, EncodeID(nodeID, actorID))
//	n, a := DecodeID(fullID)
//	require.Equal(t, n, nodeID)
//	require.Equal(t, a, actorID)
//
//	nodeID = NodeID(0xFFFF)
//	actorID = ActorID(0x7FFF)
//	fullID = FullID(0xFFFF7FFF)
//	require.Equal(t, fullID, EncodeID(nodeID, actorID))
//	n, a = DecodeID(fullID)
//	require.Equal(t, n, nodeID)
//	require.Equal(t, a, actorID)
//
//	nodeID = NodeID(0x7000)
//	actorID = ActorID(0x0001)
//	fullID = FullID(0x70000001)
//	require.Equal(t, fullID, EncodeID(nodeID, actorID))
//	n, a = DecodeID(fullID)
//	require.Equal(t, n, nodeID)
//	require.Equal(t, a, actorID)
//
//	nodeID = NodeID(0x0000)
//	actorID = ActorID(0x0001)
//	fullID = FullID(0x00000001)
//	require.Equal(t, fullID, EncodeID(nodeID, actorID))
//	n, a = DecodeID(fullID)
//	require.Equal(t, n, nodeID)
//	require.Equal(t, a, actorID)
//}
