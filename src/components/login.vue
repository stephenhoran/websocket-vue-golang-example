<template>
    <v-container>
        <v-form @submit.prevent="enterRoom">
            <v-row justify="center">
                <v-col md="6">
                    <v-card>
                        <v-toolbar color="primary">
                            <v-toolbar-title>Chatroom</v-toolbar-title>
                        </v-toolbar>
                        <v-card-text>
                            <v-row>
                                <v-text-field v-model="name" label="Nickname"></v-text-field>
                            </v-row>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn type="submit">Submit</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-col>
            </v-row>
        </v-form>
    </v-container>
</template>

<script>
    import {mapActions} from 'vuex'

    export default {
        name: "login",

        data: () => ({
            name: ""
        }),

        methods: {
            ...mapActions(['setName']),
            enterRoom() {
                this.setName(this.name)
                this.$router.push("chatroom")
            }
        }
    }
</script>

<style scoped>

</style>