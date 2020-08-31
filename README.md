我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

## globber

A small go utility for visualizing globs (and maybe regex file searches later) directly in the terminal.

![globber screenshot](https://raw.githubusercontent.com/jcquinlan/globber/master/screenshot.png)

![globber screenshot with warning](https://raw.githubusercontent.com/jcquinlan/globber/master/screenshot_warning.png)

## API

#### --glob
Default: `"*"`
Accepts any glob you'd like to test

#### --root
Default: `"."`
The starting location of the command (`glob`) will be relative to this location

#### --max-depth
Default: `20`
The maximum recursion level when traversing file structures.

#### --include-hidden-files
Default: `false`
Whether or not hidden directories and files should be considered in the visualization

#### --max-files-to-scan
Default: `200`
How many files to scan before stopping (useful to avoid scanning every file from root or something)

## Basic Usage
```bash
globber --glob=**/*.json
```
