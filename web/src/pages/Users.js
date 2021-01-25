import React, {useState, useEffect} from 'react'; 
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Table from 'react-bootstrap/Table';
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form'
import { Link } from 'react-router-dom';
import * as endpoints from '../constants/endpoints';
  
function Users() { 
  const [userFormShow, setUserFormShow] = useState(false)
  const [users, setUsers] = useState([])
  const [name, setName] = useState('')
  const [id, setId] = useState(null)
  const [editing, setEditing] = useState(false)

  useEffect(() => loadUsers(), [])

  const loadUsers = () => {
    fetch(endpoints.USERS)
      .then(res => res.json())
      .then(res => {
        if (res.err) {
          alert(res.err)
        }
        setUsers(res.users)
      })
      .catch(err => {
        console.log(err)
      })
  }

  const edit = (user) => {
    setName(user.name)
    setId(user.id)
    setUserFormShow(true)
    setEditing(true)
  }

  const newUser = () => {
    setName('')
    setId(null)
    setUserFormShow(true)
    setEditing(false)
  }

  const saveChanges = evt => {
    evt.preventDefault()
    let payload = editing?{user: {name, id}}:{name}
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    };
    fetch(editing ? endpoints.UPDATE_USER : endpoints.USERS, requestOptions)
      .then(response => response.json())
      .then(data => {
        if (data.err) {
          alert(data.err)
          return
        }
        loadUsers()
        setUserFormShow(false)
      });
  }

  const deleteUser = id => {
    const requestOptions = {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({user_id: id})
    };
    fetch(endpoints.DELETE_USER, requestOptions)
      .then(response => response.json())
      .then(data => {
        if (data.err) {
          alert(data.err)
          return
        }
        loadUsers()
      });
  }

  return <div>
    <Row>
      <Col md={9}>
        <h1>Manage Users</h1>
      </Col>
      <Col md={3}>
        <Button variant='success' className="float-right mt-4"
          onClick={() => newUser()}>New User</Button>
      </Col>
    </Row>

    <Table striped bordered hover variant="dark" className='users'>
      <tbody>
      {users && users.length
        ? users.map((user, index) => {
            return <tr key={user.id}>
              <td>
                  <Link to={`users/${user.id}`}>{user.name}</Link>
              </td>
              <td style={{ width: 200 }}>
                <Button variant='success' className="mr-5" onClick={() => edit(user)}>Edit</Button>
                <Button variant='danger' onClick={() => deleteUser(user.id)}>Delete</Button>
              </td>
            </tr>;
        })
          : <tr><td>No user found!</td></tr>
        }
        </tbody>
    </Table>
    
    <Modal
    size="sm"
    show={userFormShow}
    onHide={() => setUserFormShow(false)}
    aria-labelledby="example-modal-sizes-title-sm"
  >
    <Modal.Header closeButton>
      <Modal.Title id="example-modal-sizes-title-sm">
        {editing?'Update':'Create'} User
      </Modal.Title>
    </Modal.Header>
      <Modal.Body>
      <Form onSubmit={saveChanges}>
        <Form.Group controlId="formBasicEmail">
          <Form.Label>Name</Form.Label>
          <Form.Control type="text" placeholder="Name" value={name} onChange={evt => setName(evt.target.value) }/>
        </Form.Group>
          
        <Button variant="primary" type="submit">Submit</Button>
      </Form>

      </Modal.Body>
    </Modal>
    
    </div>
} 
  
export default Users; 