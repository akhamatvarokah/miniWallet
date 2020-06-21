<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Transfer Balance</h3>
      <form>
        <div class="form-group">
          <label for="exampleInputEmail1">Amount</label>
          <input name="amount" v-model="amount" type="number" class="form-control" id="exampleInputEmail1" aria-describedby="amount" placeholder="Enter Amount">
          <small id="amount" class="form-text text-muted">Please Input Amount</small>
        </div>
        <div class="form-group">
    <label for="exampleFormControlSelect1">Select User</label>
    <select v-el:elSelect v-model="modelSelect" class="form-control" id="exampleFormControlSelect1">
      <option v-for="user in users" :value="user.id" :key="user">{{user.username}}</option>
    </select>
  </div>
        <button v-on:click="sendAmount" type="button" class="btn btn-primary">Submit</button>
      </form>
    </header>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Transfer',
  data() {
    return {
      users: ''
    };
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    }
  },
   methods: {
    sendAmount: function () {
      // `this` inside methods points to the Vue instance
      let amount = parseInt(document.querySelector("input[name=amount]").value);
      var userID = parseInt(document.getElementById("exampleFormControlSelect1").value);
      const params = {
            amount: amount,
            user_id: userID,
        };

      if (amount > 0 && userID > 0) {
        var config = {
        headers: { 
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': 'Bearer ' + this.currentUser.token
          
          }
        };

        axios.post(this.$baseUrl + 'balance/transfer', params,config).then(response => {
          if (response.status == "200") {
            alert("sukses top Transfer");
            location.href = "home";
          }
        })
      }
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

    axios.get(this.$baseUrl + 'users/others', config).then(response => {
      this.users = response.data.response;
    })
  }
};
</script>
