import React from 'react'

import Card from 'react-bootstrap/Card'
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'

import { FaHammer, FaKey, FaTimes } from 'react-icons/fa'

const AuthForm = ({token,showForm,handleShowForm,handleToken,storeToken}) => {
    return (
        <>
        <br/>
        <Card>
            <Card.Header>Auth</Card.Header>
            <Card.Body>
                <Card.Title>Token</Card.Title>
                <Card.Text>
                {token}
                </Card.Text>
                <br/>
                <Form onSubmit={storeToken}>
                    <Form.Group className="mb-3" controlId="formBasicText">
                        <Form.Label>Token</Form.Label>
                        <Form.Control type="text" value={token} placeholder="Enter token" onChange={handleToken} />
                    </Form.Group>
                    <Button variant="primary" onClick={storeToken}>  <FaHammer /> Set Token </Button>
                    {" "}
                    { showForm ? <Button variant="secondary" onClick={handleShowForm}>  <FaTimes /> Cancel </Button> : null }
                </Form>
            </Card.Body>
        </Card>
        </>
    )
}

const Auth = () => {

    const [token,setToken] = React.useState('')
    const [showForm,setShowForm] = React.useState(false)
    const handleShowForm = () => {
        setShowForm( prev  => !prev)
    }
    const handleToken = (e) => {
        const token = e.target.value
        setToken(token)
    }
    const storeToken = () => {
        localStorage.setItem('token',token)
        setShowForm(false)
    }

    React.useEffect(() => {
        const token = localStorage.getItem('token')
        if (token) {
            setToken(token)
        }
    },[])

    if (token === '' || showForm) {
        return (
            <AuthForm {...{token,showForm,handleShowForm,handleToken,storeToken}} />
        )
    }

    return (
        <>
        <br/>
        <Button variant="primary" onClick={handleShowForm}> <FaKey /> Auth Token</Button>
        </>
    )
}
export default Auth