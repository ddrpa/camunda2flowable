# camunda2flowable(WIP)

作者目前参与的一个项目中使用了开源的 [Flowable 7](https://www.flowable.com/open-source/docs/) 作为流程引擎。

可能是因为 Flowable 正在推广自己的 SaaS 服务，市面上看起来缺乏好用的 Flowable 图形化编辑器（例如似乎不再维护的 Eclipse 插件和 flowable-ui）。

虽然用 XML 直接编写也不是不行，但是对于业务人员来说，图形化编辑器还是更直观一些（而且给图表中的对象有序地生成 ID 也是一件痛苦的事情）。

[Miranum: Camunda Modeler](https://marketplace.visualstudio.com/items?itemName=miragon-gmbh.vs-code-bpmn-modeler)
是一个不错的 Visual Studio Code 插件，虽然它是用来配置 Camunda 流程的，
不过由于这两个都符合 BPMN 2.0 规范，所以稍加处理就可以用于 Flowable。

题外话，就工具链来说，Camunda 更胜一筹啊。

## 当前问题

这是一个开发中的项目，所以：

- 产物丢弃了布局信息，不能直接渲染流程图
- 不能转回去，所以请在 Camunda 版本中做版本控制
- 作者用到什么节点 / 功能 / 属性就会加上去，所以不一定支持所有功能

## 支持节点

- 事件
  - StartEvent
  - EndEvent
  - IntermediateCatchEvent(Message)
- 活动
  - UserTask
    - 表单
      - Constraint: Required/Writeable/Readable
  - ServiceTask
- ExclusiveGateway
- SequenceFlow
  - tFormalExpression 条件表达式

## 计划

- 补充图像布局信息
- 支持其他节点
- Camunda 允许为表单项设置丰富的属性，需要提取后放在节点的 Documentation 中