<template>
  <v-container fluid fill-height>
    <v-row justify="center">
      <v-col md="8">
        <v-card>
          <v-toolbar>
            <v-toolbar-title>Websocket Chatroom</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-textarea reverse dense readonly :value="chatroomMessages"></v-textarea>
            <v-text-field label="message" v-model="message"></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn @click="sendMessage()">send</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "chatroom",

  data() {
    return {
      message: ""
    };
  },

  created() {
    window.addEventListener("keyup", this.enterMessage);
  },

  mounted() {
    this.$socket.send(this.nickName + " has joined the channel...");
  },

  computed: {
    ...mapGetters(["nickName", "chatroomMessages"])
  },

  methods: {
    enterMessage(e) {
      if (e.key === "Enter") {
        this.sendMessage();
      }
    },
    sendMessage() {
      this.$socket.send(this.nickName + ": " + this.message);
      this.message = "";
    }
  }
};
</script>

<style scoped>
</style>