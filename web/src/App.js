import React from 'react';
import { 
  BrowserRouter as Router, 
  Route, 
  Link, 
  Switch 
} from 'react-router-dom'; 

import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import Form from 'react-bootstrap/Form'
import FormControl from 'react-bootstrap/FormControl'

import Users from './pages/Users';
import Tasks from './pages/Tasks';

import './App.css';

const App = () => (
  <Container className="p-3">
  <Router> 
    <Navbar bg="light" expand="lg">
      <Navbar.Brand href="/">Task Manager</Navbar.Brand>
      <Navbar.Toggle aria-controls="basic-navbar-nav" />
      <Navbar.Collapse id="basic-navbar-nav">
        <Nav className="mr-auto">
          <Link to="/" className='nav-link'>Users</Link>
        </Nav>
        <Form inline>
          <FormControl type="text" placeholder="Find user" className="mr-sm-2" />
          <Button variant="outline-success">Search</Button>
        </Form>
      </Navbar.Collapse>
    </Navbar>

      <Switch>
        <Route path="/" component={Users} exact />
        <Route path="/users/:userId" children={<Tasks />} />
      </Switch>
    </Router> 
  </Container>
);

export default App;
