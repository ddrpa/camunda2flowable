# camunda2flowable(WIP)

作者目前参与的一个项目使用了开源的 [Flowable 7](https://www.flowable.com/open-source/docs/) 作为流程引擎。

可能是因为 Flowable 正在推广自己的 SaaS 服务，市面上看起来缺乏好用的 Flowable 图形化编辑器（例如似乎不再维护的 Eclipse 插件和 flowable-ui）。

你可能会问，直接写 XML 的流程定义不行吗？

当你开始处理稍稍复杂一些的业务流程时，就该意识到 BPMN 2.0 标准其实是一种描述业务流的 DSL/编程语言，
图形化编辑器不是为了让缺乏领域知识的人能够配置业务（虽然广告确实是这样画的饼），
而是提供了一个 Profiler 辅助查看和分析业务流，
这种情况下直接看 XML 就和汇编没什么两样（而且手写的话给图表中的对象有序生成 ID 也是件困难的事情）。

[Miranum: Camunda Modeler](https://marketplace.visualstudio.com/items?itemName=miragon-gmbh.vs-code-bpmn-modeler)
是一个用来为 Camunda 流程引擎配置业务的 Visual Studio Code 插件，
不过由于这两个都符合 BPMN 2.0 规范（而且是一个项目上 fork 出来的），所以稍加处理就可以用于 Flowable。

（就工具链来说，Camunda 似乎更胜一筹，不过论集成到项目还是 Flowable 简单些。）

## 使用方法


**前置条件**

1. 安装 Visual Studio Code 编辑器和 [Miranum: Camunda Modeler](https://marketplace.visualstudio.com/items?itemName=miragon-gmbh.vs-code-bpmn-modeler) 插件
2. 在 Visual Studio Code 中创建 `.bpmn` 文件，开始编辑流程图


`camunda2flowable -s example.bpmn` 将会读取 `example.bpmn` 的定义，生成 Flowable 流程定义并输出到 STDOUT。

可以使用管道将输出重定向到文件中，例如 `camunda2flowable -s example.bpmn > $PROJECT_DIR/resources/processes/example.flowable.bpmn`。
也可以使用 `watch` 等命令监控文件的变化。

## 附加功能

如果你像我一样使用 `<Documentation />` 节点来存储一些不被流程引擎支持/需要的信息，
可以在 Camunda 定义的同级目录创建同名文件夹，该文件夹中的 JSON 文件内容会被复制到对应的
`StartEvent` 或 `UserTask` 的 `<Documentation />` 节点中，例如：

```xml
<startEvent id="non-start-event" ...
  <documentation>
    <![CDATA[
      ${ what's in the example/non-start-event.json }
    ]]>
  </documentation>
```

流程引擎启动后可以根据相关 API 取得这部分数据。

## 当前问题

这是一个开发中的项目，所以：

- 转换产物丢弃了布局信息，不能直接渲染流程图
- 不能从 Flowable 定义转回 Camunda 定义，所以请保留 Camunda 版本的定义做版本控制
- 作者用到什么节点 / 功能 / 属性就会加上去，所以不一定支持所有功能
