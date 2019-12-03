<template>
    <v-row no-gutters>

        <v-col cols="3" v-if="this.$root.isLogin">
            <AdminEntries></AdminEntries>
        </v-col>
        <v-col cols="9" v-if="!this.$root.isLogin">
            <AdminLogin></AdminLogin>
        </v-col>
        <v-col cols="9" v-if="this.$root.isLogin && editing !== null">
            <v-card outlined>
                <v-card-text>
                    <h2>{{ editing.Id }}</h2>
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
                                rows="30"
                        ></v-textarea>
                        <v-textarea
                                v-model="editing.More"
                                label="More"
                                required
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

    export default {
        name: 'Admin',
        components: {
            AdminLogin,
            AdminEntries,
        },
        mounted() {
            axios.get('/admin/api/is_login').then(() => {
                this.$root.isLogin = true
            })
        },
        data: () => ({
            editing: null
        }),
        methods: {
            send() {
                axios.post('/admin/api/entries', this.editing).then((res) => {
                    console.log(res)
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
