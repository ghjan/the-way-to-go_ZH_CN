#前言
###***用更少的代码，更短的编译时间，创建运行更快的程序，享受更多的乐趣***

对于学习Go编程语言的爱好者来说，这本书无疑是最适合你的一本书籍，这里包含了当前最全面的学习资源。本书通过对官方的在线文档、名人博客、书籍、相关文章以及演讲的资料收集和整理，并结合我自身在软件工程、编程语言和数据库开发的授课经验，将这些零碎的知识点组织成系统化的概念和技术分类来进行讲解。

随着软件规模的不断扩大，诸多的学者和谷歌的开发者们在公司内部的软件开发过程中开始经历大量的挫折，在诸多问题上都不能给出令人满意的解决方案，尤其是在使用C++来开发大型的服务端软件时，情况更是不容乐观。由于二进制文件一般都是非常巨大的，因此需要耗费大量的时间在编译这些文件上，同时编程语言的设计思想也已经非常陈旧，这些情况都充分证明了现有的编程语言已不符合时下的生产环境。尽管硬件在过去的几十年中有了飞速的发展，但人们依旧没有找到机会去改变C++在软件开发的重要地位，并在实际开发过程中忍受着它所带来的令人头疼的一些问题。因此学者们坐下来总结出了现在生产环境与软件开发之间的主要矛盾，并尝试设计一门全新的编程语言来解决这些问题。

以下就是他们讨论得出的对编程语言的设计要求：
- 能够以更快的速度开发软件
- 开发出的软件能够很好地在现代的多核计算机上工作
- 开发出的软件能够很好地在网络环境下工作
- 使人们能够享受软件开发的过程

Go语言就在这样的环境下诞生了，它让人感觉像是Python或Ruby这样的动态语言，但却又拥有像C或者Java这类语言的高性能和安全性。

Go语言出现的目的是希望在编程领域创造最实用的方式来进行软件开发。它并不是要用奇怪的语法和晦涩难懂的概念来从根本上推翻已有的编程语言，而是建立并改善了C，Java，C#中的许多语法风格。它提倡通过接口来针对面向对象编程，通过 goroutines 和 channels 来支持并发和并行编程。

这本书是为那些想要学习Go这门全新的，迷人的和充满希望的编程语言的开发者量身定做的。当然，你在学习Go语言之前需要具备一些关于编程的基础知识和经验，并且拥有合适的学习环境，但你并不需要对C或者Java或其它类似的语言有非常深入的了解。

对于那些熟悉C或者面向对象编程语言的开发者，我们将会在本书中用Go和一些编程语言的相关概念进行比较（书中会使用大家所熟知的缩写“OO”来表示面向对象）。

本书将会从最基础的概念讲起，同时也会讨论一些类似在应用 goroutines 和 channels 时有多少种不同的模式，如何在Go语言中使用谷歌API，如何操作内存，如何在Go语言中进行程序测试和如何使用模板来开发Web应用这些高级概念和技巧。

在本书的第一部分，我们将会讨论Go语言的起源（第1章），以及如何安装Go语言（第2章）和开发环境（第3章）。

在本书的第二部分，我们将会带领你贯穿Go语言的核心思想，譬如简单与复杂类型（第4，7，8章），控制结构（第5章），函数（第6章），结构与方法（第10章）和接口（第11章）。我们会对Go语言的函数式和面向对象编程进行透彻的讲解，包括如何使用Go语言来构造大型项目（第9章）。

在本书的第三部分，你将会学习到如何处理不同格式的文件（第12章）和如何在Go语言中巧妙地使用错误处理机制（第13章）。然后我们会对Go语言中最值得称赞的设计 goroutines 和 channels 进行并发和多核应用的基本技巧的讲解（第14章）。最后，我们会讨论如何将Go语言应用到分布式和Web应用中的相关网络技巧（第15章）。

我们会在本书的第四部分向你展示许多Go语言的模式和一些默认约定，以及一些非常有用的代码片段（第18章）。在前面章节完成对所有的Go语言技巧的学习之后，你将会学习如何构造一个完整Go语言项目（第19章），然后我们会介绍一些关于Go语言在云（Google App Engine）方面的应用（第20章）。在本书的最后一章（第21章），我们会讨论一些在全世界范围内已经将Go语言投入实际开发的公司和组织。本书将会在最后给出一些对Go语言爱好者的引用，Go相关包和工具的参考，以及章节练习的答案和所有参考资源和文献的清单。

Go语言有一个被称之为“没有废物”的宗旨，就是将一切没有必要的东西都去掉，不能去掉的就无底线地简化，同时追求最大程度的自动化。他完美地诠释了敏捷编程的KISS秘诀：短小精悍！
