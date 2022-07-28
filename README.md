# 立项目的
灵感来自于边沁的Panopticon（环形监狱）理论，有关每个人向其他人展示自己包含todolist完成情况与番茄学习法时间表的主页。
Inspired by Bianchin's Panopticon theory about each person showing others their home page containing todolist completion with a tomato learning schedule.
# 技术
基础思路是Golang API服务器，连接到一个 _React_ 前端页面。

## 后端部分:
### 数据库：
- Docker-MySQL作为数据库
- dbdiagram.io生成数据库语句
- GORM作为ORM，与数据库进行交互
### 基础框架：
- 基本框架来自于Gin
- 但用 gorilla/mux 作为请求路由器Request router
### 其他：
- 用Logrus记录日志

## 前端部分：
- Sketch2Code生成前端HTML代码
- React.js调用后端API

## 其他
### 约定：
- 前端React端口：3000
- 后端传输端口：3001
- 例如http://localhost:3001/notes

# 产品设计
## 用户界面
- to do list 的状态(违约次数与完成次数)
- 番茄学习法 的状态(违约次数与完成次数，学习的时间，历次学习的经历)
- **分别对目标履行态度和学习态度的综合计算称为评分**
- **随机的其他用户的页面连接与评分**
## to do list：
- 任务ID
- 任务描述 Task description
- 任务完成状态 Completed status
	- 一个完成按钮
	- 一个违约按钮
- 违约/超时完成 次数
- 按时完成次数

## 番茄学习法：
- 使用流程：选择"番茄"个数与"番茄"之间休息的时间(3~5分钟) --> 选择每四个"番茄"间长休息的时间(15~30分钟)
- 一个"番茄"时间(25分钟)不可分割
- 四个"番茄"与之间的小休息时间加最后的大休息时间显示一个基础框架(称为篮子)，在DOM中是一个大节点Bigtomato
- 不满四个"番茄"的节点另写一个小节点Litetomato
- "番茄"个数为N时，渲染应满足 (N%4) 个大节点Bigtomato，1个小节点Litetomato
- 一个终止按钮，清除DB中最后一次记录，并且向违约部分添加记数1
- 取出DB中相关数据，渲染成网页展示历次学习的历程。
- 应该包括的状态：
	- 任务ID
	- 任务描述 Task description
	- 任务完成状态 Completed status
		- 一个完成按钮
		- 一个违约按钮
	- 违约/超时完成 次数
	- 按时完成次数




