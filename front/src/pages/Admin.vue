<template>
    <v-row no-gutters>

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
                            <v-card-text style="height: 90vh;">
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
                        <v-textarea
                                v-model="editing.Body"
                                label="Body"
                                required
                                filled
                                rows="30"
                        ></v-textarea>
                        <v-textarea
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

        </v-col>

    </v-row>
</template>

<script>
    import axios from 'axios';
    import AdminLogin from "../components/AdminLogin";
    import AdminEntries from "../components/AdminEntries";
    import Entry from "../components/Entry";

    export default {
        name: 'Admin',
        components: {
            Entry,
            AdminLogin,
            AdminEntries,
        },
        mounted() {
            axios.get('/admin/api/is_login').then(() => {
                this.$root.isLogin = true
            })
        },
        data: () => ({
            editing: null,
            dialog: false,
            ts: ''
        }),
        methods: {
            send() {
                axios.post('/admin/api/entries', this.editing).then((res) => {
                    this.ts = new Date().getTime()
                }).catch((err) => {
                    alert(err)
                })
            }
        }
    }
</script>

<style>
    .main {
        max-width: 1080px !important;
    }
</style>
