<template>
    <v-row no-gutters>

        <v-col cols="4" v-if="this.$root.isLogin">
            <AdminEntries></AdminEntries>
        </v-col>
        <v-col cols="8" v-if="!this.$root.isLogin">
            <AdminLogin></AdminLogin>
        </v-col>
        <v-col cols="8" v-if="this.$root.isLogin && editing !== null">
            <v-form
                    ref="form"

            >
                <h3>Edit</h3>
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
                <v-btn
                        color="success"
                        class="mr-4"
                        @click="send"
                >
                    Send
                </v-btn>

            </v-form>
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
            axios.get('/admin/api/is_login').then(() => {this.$root.isLogin = true})
        },
        data: () => ({
            editing: null
        }),

    }
</script>
