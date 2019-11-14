<template>
    <entry show_detail=true entry="entry"></entry>
</template>

<script>
    import axios from 'axios';
    import Entry from "../components/Entry";

    export default {
        name: 'ShowEntry',
        components: {Entry},
        data: () => ({
            entry: {},
            entryId: 0,
        }),
        mounted() {
            if (!this.entry.id) {
                this.entryId = parseInt(this.$route.params.id, 10) || 0
                axios.get('/api/entries/' + this.entryId).then(res => {
                    this.entry = res.data;
                    console.log(res.data);
                })
            }
        },

    }
</script>
