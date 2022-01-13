<template>
  <div class="navigation">


    <v-app-bar
        app
        color="light-green darken-1"
        dark
    >

      <v-toolbar-title href="/">
        <router-link tag="h1" to="/">
          <span v-if="! this.$root.isLogin">Pistatium blog</span>
          <span v-if="this.$root.isLogin">Note</span>
        </router-link>
      </v-toolbar-title>

      <v-spacer/>
      Menu
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app right dark class="dark-green darken-4" style="padding-top: 12px;">
      <v-list dense>
        <v-list-item link to="/">
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title>Top</v-list-item-title>
            <v-list-item-subtitle>Back to blog top.</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-divider style="margin: 4px 0 12px 0"></v-divider>

        <v-subheader style="height: 20px">About Blog</v-subheader>
        <v-list-item>
          <v-list-item-content class="about_blog">
            プログラミングの小ネタとか、買ったもののレビューとかいろいろ。
            このブログは Google App Engine を利用した自作システムで動いています。
          </v-list-item-content>
        </v-list-item>


        <!--                <v-list-item link to="/portfolio">-->
        <!--                    <v-list-item-action>-->
        <!--                        <v-icon>mdi-notebook-outline</v-icon>-->
        <!--                    </v-list-item-action>-->
        <!--                    <v-list-item-content>-->
        <!--                        <v-list-item-title>Portfolio</v-list-item-title>-->
        <!--                        <v-list-item-subtitle>About works.</v-list-item-subtitle>-->
        <!--                    </v-list-item-content>-->
        <!--                </v-list-item>-->

        <v-divider style="margin: 4px 0 12px 0"></v-divider>

        <v-subheader style="height: 20px">Author</v-subheader>
        <v-list-item>
          <v-list-item-avatar>
            <v-img src="/img/profile.png"
                   style="border-radius: 32px; border: #ffffff 1px solid;background: #ffffff"></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>kimhiro_n</v-list-item-title>
            <v-list-item-subtitle>
              <a href="https://twitter.com/kimihiro_n" target="_blank" rel="noopener">Twitter</a> /
              <a href="https://github.com/pistatium" target="_blank" rel="noopener">GitHub</a> /
              <a href="https://qiita.com/kimihiro_n" target="_blank" rel="noopener">Qiita</a>
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-divider style="margin: 4px 0 12px 0"></v-divider>

        <v-subheader style="height: 20px">Apps</v-subheader>

        <v-list-item link href="https://emiru.pistatium.dev/"
                     target="_blank" rel="noopener">
          <v-list-item-action>
            <v-img style="width: 32px" src="https://emiru.pistatium.dev/images/icon-128x128.png"></v-img>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Emiru</v-list-item-title>
            <v-list-item-subtitle>
              フォロー中の絵師さんの<br>絵を漏らさず見るための<br>専用Viewer
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link href="https://emiru.pistatium.dev/"
                     target="_blank" rel="noopener">
          <v-list-item-action>
            <v-img style="width: 32px"
                   src="https://raw.githubusercontent.com/pistatium/mahougen/master/app/src/main/res/mipmap-xxxhdpi/icon192.png"></v-img>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Mahougen</v-list-item-title>
            <v-list-item-subtitle>
              魔方陣をサクッと生成<br>なぞるだけで楽しい<br>Androidアプリ
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-list-item link href="https://play.google.com/store/apps/dev?id=8846644054781474049&hl=ja"
                     target="_blank" rel="noopener">
          <v-list-item-action>
            <v-icon style="width:32px">mdi-android</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>その他のアプリ</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: "GlobalNavigation",
  data: () => ({
    drawer: null,
  }),
}
</script>

<style scoped>
h1 {
  cursor: pointer;
  font-size: 24px;
}

.navigation {
  margin-bottom: 16px;
}

.about_blog {
  font-size: 80%;
  color: #9f9f9f;
}

.site-icon {
  float: left;
  margin-right: 6px;
}

</style>
