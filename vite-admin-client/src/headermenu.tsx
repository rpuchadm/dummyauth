import React from 'react'

import Container from 'react-bootstrap/Container'
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'

const HeaderMenu: React.FC = () => {
  return (
    <Navbar bg="primary" data-bs-theme="light">
    <Container>
    <Navbar.Brand href="#home">Auth Admin</Navbar.Brand>
    <Nav className="me-auto">
        <Nav.Link href="#home">Home</Nav.Link>
        <Nav.Link href="#features">Features</Nav.Link>
        <Nav.Link href="#pricing">Pricing</Nav.Link>
    </Nav>
    </Container>
</Navbar>
  )
}

export default HeaderMenu