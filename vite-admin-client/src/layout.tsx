import React from 'react'

import Container from 'react-bootstrap/Container'
import Col from 'react-bootstrap/Col'
import Row from 'react-bootstrap/Row'

const Layout: React.FC = ({ children }) => {
  return <Container><Row><Col>{children}</Col></Row></Container>
}

export default Layout