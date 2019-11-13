<template>
    <entry show_detail=true entry="entry"></entry>
</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";

    export default {
        name: 'Entry',
        components: {Entry},
        props: ['entry'],
        data: () => ({
            entry: {},
            page: 0,
        }),
        mounted() {
            if (!entry) {
                alert('')
            }
        },

    }
</script>
