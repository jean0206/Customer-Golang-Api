<template>
  <div class="home">
    <div class="title">
      Clientes
    </div>

    <div v-if="customers.length!=0" class="customer-content">

    <div  class="customers" v-for="(item,key) in customers" :key="key">
      <UserCard :name=item.name  @clicked="getInfo(key)" :id=item.id :age=item.age></UserCard>
    </div>
    </div>

    <div class="not-found" v-else> 
      No hay información de los clientes
      <img src="../assets/not_found.png"/>
    </div>
    
  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios'
import UserCard from '../components/CardUser'
export default {
  name: 'Home',
  components: {
    UserCard
  },
  data: function(){
    return {
      customers:[]
    }
  },
  methods: {
    getInfo(index){
      console.log(typeof(this.customers[index]))
      localStorage.setItem("customerInfo",JSON.stringify(this.customers[index]))
      this.$router.push("/about")
    }
  },
  async beforeMount() {
    await axios.get('http://localhost:3000/customer/').then(response=>{
     this.customers=(response.data['customers'])
      console.log(this.customers)
    }).catch(error=>{
      console.log(error.message)
    })
    

  },
}
</script>

<style scoped>
.home{
  display: flex;
  flex-direction: column;
  text-align: center;

  margin-top: 50px;
  align-items: center;
}

.title{
    font-size: 40px;
}

.customers{
  margin-top: 30px;
  margin-left: 20px;

}
.customer-content{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.not-found{
  width: 500px;
  height: 500px;
  display: flex;
  flex-direction: column;
  margin-top: 50px;
}
</style>
