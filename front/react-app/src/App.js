import { useState } from 'react'

const App = () => {
  const [persons, setPersons] = useState([
    { name: 'Arto Hellas', number: '040-123456', id: 1 },
    { name: 'Ada Lovelace', number: '39-44-5323523', id: 2 },
    { name: 'Dan Abramov', number: '12-43-234345', id: 3 },
    { name: 'Mary Poppendieck', number: '39-23-6423122', id: 4 }

  ])
  const [newName, setNewName] = useState('write name')
  const [newNumber, setNewNumber] = useState('write number')
  const changeperson =(event)=>{
    event.preventDefault()
    console.log('button clicked', event.target)
    const personObject = {
      name: 'write name' ,
      number : 'write number'
    }
  
    setPersons(persons.concat(personObject)) 
    setNewName('')
    setNewNumber('')
  } 
  const handleNameChange = (event) => {   
    console.log(event.target.value)    
    setNewName(event.target.value)     
    }
  const handleNumberChange = (event) => {   
      console.log(event.target.value)    
      setNewNumber(event.target.value)     
    }
  return (
    <div>
      <h2>Phonebook</h2>
      <form>
        <div>
        Name: <input value={newName} onChange={handleNameChange}/>
        </div>
        <div>
          Number: <input value={newNumber} onChange={handleNumberChange}/>
          </div>
        <div>
        <button type="submit" onClick={changeperson}>add</button>
        </div>
      </form>
      <h2>Numbers</h2>
      <div>
      {persons.map(person => <ul>{person.name}  :  {person.number}</ul>)}
      </div>
    </div>
  )
}

export default App