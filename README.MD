# AI Magic Music

基于 NextJS 开发的音乐生成网站，可以通过用户输入的 prompt 在一分钟左右的时间生成想要的歌曲，也可以通过通过上传个人语音和音乐片段，实现 AI 音乐翻唱。

## 快速开始

1. 获取 app.suno.ai 账户的 cookie 

找到包含关键词 "client?_clerk_js_version" 的请求。找到请求的 Cookie 部分，并复制 Cookie 的值

2. 克隆项目

```shell
git clone https://github.com/YellowPancake/AI-Magic-Music.git
```

3. 安装依赖

```shell
cd AI-Magic-Music
pnpm install
```

4. 初始化数据库

使用本地数据库： [local postgres](https://wiki.postgresql.org/wiki/Homebrew)
或者使用在线数据库： [vercel-postgres](https://vercel.com/docs/storage/vercel-postgres)

在 `data/install.sql` 文件中复制创建数据库用到的 sql

5. 设置环境变量

在项目跟目录添加 `.env.local` 文件，填入如下配置：

```
NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=""
NEXT_PUBLIC_CLERK_SIGN_IN_URL=/sign-in
NEXT_PUBLIC_CLERK_SIGN_UP_URL=/sign-up
NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL=/
NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL=/

SUNO_COOKIE=""

LEMON_SQUEEZY_HOST=https://api.lemonsqueezy.com/v1
LEMON_SQUEEZY_API_KEY=
LEMON_SQUEEZY_STORE_ID=
LEMON_SQUEEZY_PRODUCT_ID=
LEMON_SQUEEZY_MEMBERSHIP_MONTHLY_VARIANT_ID=
LEMON_SQUEEZY_MEMBERSHIP_SINGLE_TIME_VARIANT_ID=
LEMONS_SQUEEZY_SIGNATURE_SECRET=

POSTGRES_URL=
```

SUNO_COOKIE 是你第一步获取的 cookie 值

6. 本地开发

```shell
pnpm dev
```

打开预览：`http://localhost:3000` 

## 致谢
- [Suno AI](https://www.suno.ai)
- [nextjs](https://nextjs.org/docs)
- [clerk](https://clerk.com/docs/quickstarts/nextjs)
- [node-postgres](https://node-postgres.com/)
- [tailwindcss](https://tailwindcss.com/)
- [Lemon Squeezy](https://www.lemonsqueezy.com/)
- [aiwallpaper](https://github.com/all-in-aigc/aiwallpaper)
- [suno-music-generator](github.com/Alvin-Liu/suno-music-generator)

## ToDo list
- [ ] Vercel 一键部署
- [ ] AI 音乐翻唱加入网站界面