import { useState } from 'react'
import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';
import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import Cashback from './Receipt'

function CashBackForm() {
    const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);
  return (
    <>
    <Form>
        <Container className="w-50 mt-3 pt-3">
            <h3> Create a Free Cashback Withdraw Point </h3>
      <Form.Group className="mb-3" controlId="formBasicEmail">
        <Form.Label>Store Name </Form.Label>
        <Form.Control type="text" placeholder="Store Name" />
        
      </Form.Group>

      <Form.Group className="mb-3" controlId="formBasicPassword">
        <Form.Label>Cashback Amount</Form.Label>
        <Form.Control type="text" placeholder="$0.0" />
      </Form.Group>
      <Button variant="secondary" onClick={handleShow}>
        Create Cashback
      </Button>
      </Container>
    </Form>

<Modal show={show} onHide={handleClose} 
size="md"
fullscreen={true}
className="me-2 mb-2"
>
<Modal.Header closeButton>
  <Modal.Title>Thanks for visiting the Fitness Store! </Modal.Title>
</Modal.Header>
<Modal.Body className="text-center">
    <p className="p-4"> Here is your 100 Satoshi CashBack!!</p>
    <p className="p-4">Scan this QR Code To Claim it!!</p>
    <Cashback />
</Modal.Body>

</Modal>
</>
  );
}

export default CashBackForm;