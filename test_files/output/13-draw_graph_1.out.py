import networkx as nx
import matplotlib.pyplot as plt
G = nx.Graph()
G.add_edge("Kenichi", "Dini")
G.add_edge("Kenichi", "Rodrigo")
G.add_edge("Dini", "Rodrigo")
nx.draw(G, with_labels=True, font_weight='bold')
plt.show()
