import Vue from 'vue';
import Vuetify from 'vuetify/lib';
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);

export default new Vuetify({ theme: {
    themes: {
        light: {
            primary: colors.lightGreen.lighten1,
                secondary: colors.green.lighten4,
                accent: colors.green.darken2,
        },
    },
}});
