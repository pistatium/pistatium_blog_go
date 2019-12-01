<template>
    <v-form
            ref="form"
            lazy-validation
    >
        <h3>Login</h3>
        <v-text-field
                v-model="username"
                label="Username"
                required
        >

        </v-text-field>
        <v-text-field
                v-model="password"
                label="Password"
                type="password"
                required
        ></v-text-field>
        <v-btn
                color="success"
                class="mr-4"
                @click="login"
        >
            Login
        </v-btn>

    </v-form>
</template>

<script>
    import axios from 'axios';
    export default {
        name: "AdminLogin",
        data: () => ({
            username: "",
            password: "",
            session: "",
        }),
        mounted: function () {
        },
        methods: {
            login() {
                axios.post("/admin/login", {
                    username: this.username,
                    password: this.password
                }).then(res => {
                    this.$root.isLogin = true;
                    console.log(this.$root)
                })
            }
        }
    }
</script>

<style scoped>

</style>
