graph = TinkerGraph.open()
g = traversal().withEmbedded(graph)

cnt = 1

g.addV('user').property('name', 'ユーザー1').property(id, cnt++).next()

for (i in 1..1e4) {
  g.addV('post')
  .property('name', '投稿' + i)
  .property(id, cnt++)
  .next()
}

for (i in 1..1e4) {
  g.addE('post').from(g.V(1).next()).to(g.V(i + 1).next()).property(id, cnt++).next()
}

println clockWithResult(1) { g.V().has('user', 'name', 'ユーザー1').out().hasLabel('post').values('name') }