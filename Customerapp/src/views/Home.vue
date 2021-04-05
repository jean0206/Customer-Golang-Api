<template>
  <div class="home">


      <NavBar @clicked="openModal"></NavBar>
    <div v-if="modal" class="modal-container">
      <div class="modal">
        <div class="close-button">
     

          <font-awesome-icon
          icon="times-circle"
          @click="openModal"
          :style="{ color: '#424AF5', paddingRight:20,paddingTop:10}"
          size="2x"
        />
   
        </div>
        <input v-model="date" class="date-input" type="date"/>
        <button @click="getDate">Importar </button>
    </div>
    </div>
    

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
import NavBar from '../components/NavBar'
export default {
  name: 'Home',
  components: {
    UserCard,
    NavBar
  },
  data: function(){
    return {
      customers:[],
      modal:false,
      date:null
    }
  },
  methods: {
    async getDate(){
     await  axios.post("http://localhost:8080/alter",{
        "drop_op": "DATA"
      }).then(response=>{
        console.log(response.data)
      })
      console.log(typeof(this.date))
      const datum = Date.parse(this.date);
      const newDate = datum/1000
      await axios.get('http://localhost:3000/data/?date='+newDate).then(response=>{

    }).catch(error=>{
      console.log(error.message)
    })
    this.getCustomers()

    },
    openModal(){
      this.modal=!this.modal
    },
    getInfo(index){
      console.log(typeof(this.customers[index]))
      localStorage.setItem("customerInfo",JSON.stringify(this.customers[index]))
      this.$router.push("/about")
    },
    async getCustomers(){
      this.customers=[]
      await axios.get('http://localhost:3000/customer/').then(response=>{
     this.customers=(response.data['customers'])
      console.log(this.customers)
    }).catch(error=>{
      console.log(error.message)
    })
    }
  },
  async beforeMount() {
     await this.getCustomers()    

  },
}
</script>

<style scoped>
.home{
  display: flex;
  flex-direction: column;
  text-align: center;


  align-items: center;
}

.title{
    font-size: 40px;
    margin-top: 60px;
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

.close-button{
  display: flex;
  justify-content: flex-end;
  width: 100%;
}

.date-input{
  width: 250px;
}

.modal{
  width: 400px;
  height: 400px;
  background-color: white;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
}

.modal button{
  width: 150px;
  height: 40px;
  border: none;
  outline: none;
  border-radius: 10px;
  font-family: 'Roboto';
  font-weight: bold;
  color: white;
  background-color: #2e32ac;;
  cursor: pointer;
  margin-bottom: 40px;
}

.modal button:hover{
  background-color: #2e32acaf;
}

.modal-container{
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: center;
  position: fixed;
  height: 100%;
  align-items: center;
  background-color: rgba(161, 161, 161,0.4);
  margin-top: -6px;
}

.nav{
  width: 100%;
}
</style>
