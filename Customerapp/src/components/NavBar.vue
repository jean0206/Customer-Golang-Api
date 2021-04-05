<template>
    <nav class="nav-bar">
        <button @click="clicked">Importar <font-awesome-icon
          icon="arrow-alt-circle-up"
          :style="{ color: '#424AF5' }"
          size="2x"
        /></button>
    </nav>
</template>

<script>

export default ({
    data: ()=>{
        return{

        }
    },methods: {
        clicked(){
          this.$emit('clicked')
      }
    },
})
</script>

<style scoped>
.nav-bar{
    width: 100%;
    height: 60px;
    background-color: #424AF5;
    margin-top: -8px;
     display: flex;
    justify-content: flex-end;
    align-items: center;
   
}

.nav-bar button{
    outline: none;
    border: none;
    height: 35px;
    display: flex;
    align-items: center;
    margin-right: 20px;
    border-radius: 16px;
    font-family: 'Roboto';
    font-weight: 600;
    justify-content: space-between;
    background-color: aliceblue;
    cursor: pointer;
}

.nav-bar button:hover{
    background-color: #bdbed3;
}
</style>
