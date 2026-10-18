/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-18 16:27:56
 * @link    github.com/taseikyo
 */

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 初始化入度表和邻接表
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	for _, pre := range prerequisites {
		// pre = [a, b], 表示 先修 b 才能修 a
		// 构建邻接表: b -> a
		graph[pre[1]] = append(graph[pre[1]], pre[0])
		// 课程 a 的入度加1
		inDegree[pre[0]]++
	}

	// 2. 将所有入度为0的课程加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. 记录已修完的课程数量
	finished := 0

	// 4. BFS 拓扑排序
	for len(queue) > 0 {
		// 取出队首课程
		course := queue[0]
		queue = queue[1:]
		finished++

		// 遍历所有需要以 course 为先修课的后续课程
		for _, next := range graph[course] {
			inDegree[next]-- // 依赖的先修课少了一门
			if inDegree[next] == 0 {
				queue = append(queue, next) // 新的可修课程
			}
		}
	}

	// 5. 判断是否所有课程都已修完
	return finished == numCourses
}
