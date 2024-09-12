import networkx as nx
import matplotlib.pyplot as plt
G = nx.Graph()
edge_labels={}
G.add_edge("rodrigo", "kenichi")
edge_labels[("rodrigo", "kenichi")] = "friends_with"
G.add_edge("rodrigo", "dini")
edge_labels[("rodrigo", "dini")] = "friends_with"
G.add_edge("rodrigo", "colega")
edge_labels[("rodrigo", "colega")] = "studies_with"
pos = nx.spring_layout(G)
nx.draw(G, pos, node_color='pink', alpha=0.9, node_size=500, labels={node: node for node in G.nodes()})
nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, font_color='red')
plt.show()
