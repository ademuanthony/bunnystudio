import React, {useState, useEffect} from 'react'; 
import { useParams } from "react-router";
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col'; 
import Table from 'react-bootstrap/Table'; 
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import * as endpoints from '../constants/endpoints';
  
function Tasks() { 
  let { userId } = useParams()
  const [user, setUser] = useState({})
  const [tasks, setTasks] = useState([])
  const [description, setDescription] = useState('')
  const [id, setId] = useState(null)
  const [state, setState] = useState('to do')
  const [userFormShow, setUserFormShow] = useState(false)
  const [editing, setEditing] = useState(false)

  useEffect(() => loadDetails(), [])

  const loadDetails = () => {
    loadUser()
    loadTasks()
  }

  const loadUser = () => {
    fetch(endpoints.USER_DETAILS + `?userid=${userId}`)
      .then(res => res.json())
      .then(res => {
        if (res.err) {
          alert(res.err)
          return
        }
        setUser(res.user)
      })
      .catch(err => {
        console.log(err)
      })
  }

  const loadTasks = () => {
    fetch(endpoints.TASKS + `?userid=${userId}`)
      .then(res => res.json())
      .then(res => {
        if (res.err) {
          alert(res.err)
        }
        setTasks(res.tasks)
      })
      .catch(err => {
        console.log(err)
      })
  }

  const edit = (task) => {
    setDescription(task.description)
    setId(task.id)
    setState(task.state)
    setUserFormShow(true)
    setEditing(true)
  }

  const newTask = () => {
    setDescription('')
    setId(null)
    setState('to do')
    setUserFormShow(true)
    setEditing(false)
  }

  const saveChanges = evt => {
    evt.preventDefault()
    let model = { description, state, user_id: userId }
    if (editing) {
      model.id = id
      model = {task: model}
    }
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(model)
    };
    fetch(editing ? endpoints.UPDATE_TASK : endpoints.NEW_TASK, requestOptions)
      .then(response => response.json())
      .then(data => {
        if (data.err) {
          alert(data.err)
          return
        }
        loadTasks()
        setUserFormShow(false)
      });
  }

  return <div> 
     <Row>
      <Col md={9}>
        <h1>{user.name}</h1>
      </Col>
      <Col md={3}>
        <Button variant='primary' className="float-right mt-4" onClick={() => newTask()}>Add Task</Button>
      </Col>
    </Row>
    
    <Table striped bordered hover variant="dark" className='users'>
      <tbody>
      {tasks && tasks.length
        ? tasks.map((task, index) => {
            return <tr key={task.id}>
              <td>
                {task.description}
              </td>
              <td style={{ width: 200 }}>
                {task.state}
              </td>
              <td style={{ width: 200 }}>
                <Button variant='success' className="mr-5" onClick={() => edit(task)}>Edit</Button>
              </td>
            </tr>;
        })
          : <tr><td>No tasks found!</td></tr>
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
          
          <Form.Group>
            <Form.Label>Description</Form.Label>
            <Form.Control type="text" placeholder="Task description" value={description}
              onChange={evt => setDescription(evt.target.value)} />
          </Form.Group>
          
          <Form.Group>
            <Form.Label>State</Form.Label>
            <Form.Control
              value={state}
              onChange={evt => setState(evt.target.value)}
              as="select"
              className="mr-sm-2"
              id="inlineFormCustomSelect"
              custom
            >
              <option value="to do">To Do</option>
              <option value="done">Done</option>
            </Form.Control>
          </Form.Group>
          
          
        <Button variant="primary" type="submit">Submit</Button>
      </Form>

      </Modal.Body>
    </Modal>
    
    </div> 
} 
export default Tasks; 