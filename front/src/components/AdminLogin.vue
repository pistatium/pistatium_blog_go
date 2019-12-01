<template>
    <v-form
            ref="form"
            v-model="valid"
            lazy-validation
    >
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

    </v-form>
</template>

<script>

    export default {
        name: "AdminLogin"
    }
</script>

<style scoped>

</style>
