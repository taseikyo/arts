# -*- coding: utf-8 -*-

from pygraph.classes.digraph import digraph
import sqlite3

class PRIterator:
    __doc__ = '''计算一张图中的PR值'''

    def __init__(self, dg):
        self.damping_factor = 0.85  # 阻尼系数,即α
        self.max_iterations = 1000  # 最大迭代次数
        self.min_delta = 0.00001  # 确定迭代是否结束的参数,即ϵ
        self.graph = dg

    def page_rank(self):
        # 先将图中没有出链的节点改为对所有节点都有出链
        for node in self.graph.nodes():
            if len(self.graph.neighbors(node)) == 0:
                for node2 in self.graph.nodes():
                    digraph.add_edge(self.graph, (node, node2))

        nodes = self.graph.nodes()
        graph_size = len(nodes)

        if graph_size == 0:
            return {}
        # 给每个节点赋予初始的PR值，第一轮的PR值是均等的，即 1/N
        page_rank = dict.fromkeys(nodes, 1.0 / graph_size)
        # 公式中的(1−α)/N部分
        damping_value = (1.0 - self.damping_factor) / graph_size

        flag = False
        for i in range(self.max_iterations):
            change = 0
            for node in nodes:
                rank = 0
                # 遍历所有“入射”的页面
                for incident_page in self.graph.incidents(node):
                    # "入射"页面的权重根据其出链个数均分，然后传递给当前页面
                    rank += self.damping_factor * (page_rank[incident_page] / len(self.graph.neighbors(incident_page)))
                # 增加随机概率转移矩阵的部分
                rank += damping_value
                change += abs(page_rank[node] - rank)  # 绝对值
                page_rank[node] = rank

            print("This is NO.%s iteration" % (i + 1))
            print(page_rank)

            if change < self.min_delta:
                flag = True
                break
        if flag:
            print("finished in %s iterations!" % node)
        else:
            print("finished out of 100 iterations!")
        return page_rank



if __name__ == '__main__':
    # 创建一个网络拓朴图
    dg = digraph()

    conn = sqlite3.connect('zhihu.db')
    c = conn.cursor()

    nodes = []
    cursor = c.execute("SELECT DISTINCT user_url, followee_url FROM Following;")
    for row in cursor:
        #print row
        if row[0] not in nodes:
            nodes.append(row[0])
        if row[1] not in nodes:
            nodes.append(row[1])

    # 添加实体节点
    dg.add_nodes(nodes)

    cursor = c.execute("SELECT DISTINCT user_url, followee_url FROM Following;")
    for row in cursor:
        user_url = str(row[0])
        followee_url = str(row[1])
        # 添加实体间link(边)      followee_url -> user_url
        #print "followee_url:{0} -> user_url:{1}".format(followee_url, user_url)
        dg.add_edge((followee_url, user_url))

    conn.close()

    # PRrank迭代计算
    pr = PRIterator(dg)
    page_ranks = pr.page_rank()

    with open("page_ranks.txt", 'w') as fp:
        fp.write(str(page_ranks))

    print("The final page rank is\n", page_ranks)
