<template>
    <v-row no-gutters class="editor">

        <v-col cols="3" v-if="this.$root.isLogin">
            <AdminEntries v-bind:ts="this.ts"></AdminEntries>
        </v-col>
        <v-col cols="9" v-if="!this.$root.isLogin">
            <AdminLogin></AdminLogin>
        </v-col>
        <v-col cols="9" v-if="this.$root.isLogin && editing !== null">
            <v-card outlined>

                <v-card-text>
                    <h2>{{ editing.Id }}</h2>

                    <v-dialog
                            v-model="dialog"
                            width="800"
                            scrollable
                    >
                        <template v-slot:activator="{ on }">
                            <v-btn
                                    color="red lighten-2"
                                    dark
                                    v-on="on"
                            >
                                Preview
                            </v-btn>
                        </template>
                        <v-card>
                            <v-card-text style="height: 90vh;" v-if="dialog">
                                <Entry v-bind:entry="this.editing" v-bind:show_detail="true" v-bind:index="1"></Entry>
                            </v-card-text>
                        </v-card>
                    </v-dialog>

                    <v-btn
                            color="success"
                            class="mr-4"
                            @click="this.send"
                    >
                        Send
                    </v-btn>

                    <v-divider></v-divider>
                    <v-form
                            ref="form"

                    >
                        <v-text-field
                                v-model="editing.Title"
                                label="title"
                                required
                        >
                        </v-text-field>
                        <v-text-field
                                v-model="editing.Thumbnail"
                                label="thumbnail"
                                required
                        >
                        </v-text-field>
                        <v-textarea
                                id="edit_body"
                                v-model="editing.Body"
                                label="Body"
                                required
                                filled
                                rows="30"

                        ></v-textarea>
                        <v-textarea
                                id="edit_more"
                                v-model="editing.More"
                                label="More"
                                required
                                filled
                                rows="30"

                        ></v-textarea>
                        <v-switch v-model="editing.Public" label="Public"></v-switch>
                        <v-btn
                                color="success"
                                class="mr-4"
                                @click="this.send"
                        >
                            Send
                        </v-btn>
                    </v-form>
                </v-card-text>
            </v-card>
            <v-card>
                <v-card-title>画像アップロード</v-card-title>
                <v-card-actions>
                    <v-form >
                        <v-file-input name="file" accept="image/*" label="File input" v-model="file" @change="submitImage"></v-file-input>
                    </v-form>
                </v-card-actions>
                <v-card-text>
                    <div class="images" v-for="img in this.images" v-bind:key="img">
                        <img v-bind:src="img">
                        <div>
                          URL: <input readonly v-bind:value="absPath(img)" type="text" class="pathbox" onclick="this.select();">

                        </div>
                        <div>
                          Tag: <input readonly v-bind:value="imgTag(img)" type="text" class="pathbox" onclick="this.select();">
                        </div>
                    </div>
                </v-card-text>

            </v-card>
        </v-col>

    </v-row>
</template>

<script>
    import axios from 'axios';
    import AdminLogin from "../components/AdminLogin";
    import AdminEntries from "../components/AdminEntries";
    import Entry from "../components/Entry";

    export default {
        name: 'AdminConsole',
        components: {
            Entry,
            AdminLogin,
            AdminEntries,
        },
        mounted() {
            axios.get('/admin/api/is_login').then(() => {
                this.$root.isLogin = true
            })
            this.timer = setInterval(this.send,30000)
        },
        complete: function() {
            clearInterval(this.timer)
        },
        data: () => ({
            editing: null,
            dialog: false,
            ts: '',
            timer: null,
            file: null,
            images: []
        }),
        methods: {
            send() {
                if (! this.editing || ! this.editing.Id) {
                    return
                }
                axios.post('/admin/api/entries', this.editing).then((res) => {
                    this.ts = new Date().getTime()
                }).catch((err) => {
                    alert(err)
                })
            },
            submitImage() {
                let formData = new FormData();
                formData.append('file', this.file, this.file.name);
                let config = {
                    headers: {
                        'content-type': 'multipart/form-data'
                    }
                };
                axios.post('/admin/api/photos', formData, config).then((res) => {
                    console.log(res)
                    this.images.unshift(res.data.path)
                }).catch((err) => {
                    alert(err)
                })
            },
            absPath(path) {
                let port = window.location.port
                if (port !== '') {
                  port = ':' + port
                }
                return `${window.location.protocol}//${window.location.hostname}${port}${path}`
            },
            imgTag(path) {
              return `<img src="${this.absPath(path)}" loading="lazy">`
            }
            // IMEが暴発するのでPEND
            // input_tab(e) {
            //
            //     if (e.key === "Tab") {
            //         e.preventDefault();
            //         var elem = e.target;
            //         var val = elem.value;
            //         var pos = elem.selectionStart;
            //         elem.value = val.substr(0, pos) + '\t' + val.substr(pos, val.length);
            //         elem.setSelectionRange(pos + 1, pos + 1);
            //     }
            // }
        }
    }
</script>

<style scoped>
    .images {
    }

    .editor {
        margin-top: 24px;
    }
    .images img {
        max-width: 100px;
        max-height: 100px;
    }
    .pathbox {
        width: 100%;
    }

</style>
