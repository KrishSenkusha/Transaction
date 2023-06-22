import React, {useState } from 'react'; //See on utube
import { View, StyleSheet, TextInput, Text, Button } from 'react-native';

export default function App() {
  
  const [lender, setLender] = useState(""); //constant declared, function name and default value which is use state.
  const [reciever, setReciever] = useState(""); //useState is used to define a state variable and has a default value in it.
  const [date, setDate] = useState("");
  const [amount, setAmount] = useState("");

  const AddData = async () => { 
    console.warn({ lender, reciever, amount, date }); //it is used to check whether the input has come or not
    let data = {
      lender, // it is used to store the data and values to be sent to backend right 
      reciever,
      amount: parseInt(amount,10),
      date,
    };
  
      
    let result = await fetch("http://localhost:9000/Payment", { //Fetch is used to fetch the url and it is used to send the result in json format... await will wait for the function to execute
      method: "POST",
      headers: { "Content-Type": "application/json" }, // says we r sending data in json format
      body: JSON.stringify(data), //this will convert the data object to json format and send the data
    });

    result = await result.json(); // it will convert the result into json format and store it in result variable,
    console.warn(result); // it is used to check wheter my api has been called or not and print result
  };

  //View it is used to contain all the buttons container etc 
  return (
    <View style={styles.container}>  
    <Text>Lender name : </Text>
     <TextInput style={styles.input} placeholder='From'
     value={lender}
     onChangeText={(text) => setLender(text)} // it will keep the updated text stored in a variable also called state variable
     />

    <Text>Reciever Name : </Text>
     <TextInput style={styles.input} placeholder='To'
     value={reciever}
     onChangeText={(text) => setReciever(text)}
     />

    <Text>Amount : </Text>
     <TextInput style={styles.input} placeholder='₹ '
     value={amount}
     onChangeText={(text) => setAmount(text)}
     />

    <Text> Date :</Text>
     <TextInput type style={styles.input} placeholder='YYYY/MM/DD'
     value={date}
     onChangeText={(text) => setDate(text)}
     />


     <Button 
      color={'red'}
      onPress={AddData}
      title='Submit'/>
      </View>
  );
}

const styles = StyleSheet.create({
  container: {
    
    flex:1,   // It will change width according to diff devices like pc mobile etc
    backgroundColor: 'white',
    alignItems:"center",
    justifyContent:'center',  
    padding:60,
    margin:10,
    marginTop:30,
    borderRadius:30  // border radius is the curve on my edge of the textbox 
    
    
  },
  input:{   //textbox
    marginRight:20,
    padding:10,
    width:300,
    borderWidth:1, //thickness of textbox
    height:40,
    borderRadius:10,
    margin:10
 }
  
});