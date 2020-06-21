<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Balance History</h3>
      <table class="table">
        <tr>
          <td>No</td>
          <td>Name</td>
          <td>Type</td>
          <td>balance Before</td>
          <td>balance After</td>
          <td>Created</td>
        </tr>
        <tr v-for="(item, index) in results" :key="item">
            <td>{{ index + 1 }}</td>
            <td>{{ item.user_balance.user.username }}</td>
            <td>{{ item.type }}</td>
            <td>{{ item.balance_before }}</td>
            <td>{{ item.balance_after }}</td>
            <td>{{ item.created_at }}</td>
        </tr>

      </table>
    </header>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Home',
  data() {
    return {
      results: ''
    };
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    }
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }

    var config = {
      headers: { 
        'Content-Type': 'application/x-www-form-urlencoded',
        'Authorization': 'Bearer ' + this.currentUser.token
      
      }
    };

    axios.get(this.$baseUrl + 'balance/history', config).then(response => {
      this.results = response.data.response;
    })
  }
};
</script>
