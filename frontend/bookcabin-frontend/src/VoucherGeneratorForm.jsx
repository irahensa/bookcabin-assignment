import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import { Container } from "react-bootstrap";

function VoucherGeneratorForm() {
  const submitForm = (e) => {
    e.preventDefault();

    const formData = new FormData(e.target);
    const payload = Object.fromEntries(formData);

    console.log(payload);

    console.log("try hit check");
    fetch("http://127.0.0.1:8080/api/check", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },

      body: JSON.stringify(payload),
    })
      .then((response) => response.json())
      .then((data) => {
        if (!data.exists) {
          console.log("not exist");
          fetch("http://127.0.0.1:8080/api/generate", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },

            body: JSON.stringify(payload),
          })
            .then((response) => response.json())
            .then((data) => console.log(data))
            .catch((error) => console.error("Error:", error));
        } else {
          console.log("exist");
        }
      })
      .catch((error) => console.error("Error:", error));
  };

  const logvar = (e) => {
    e.preventDefault();

    console.log(checkData);
  };

  return (
    <>
      <Container id="main-container">
        <h2>Voucher Generator</h2>
        <Form id="voucher-generator-form" onSubmit={submitForm}>
          <Form.Group className="mb-3">
            <Form.Label>Crew name</Form.Label>
            <Form.Control type="text" name="name" placeholder="name" />
          </Form.Group>

          <Form.Group className="mb-3">
            <Form.Label>Crew ID</Form.Label>
            <Form.Control type="text" name="id" placeholder="id" />
          </Form.Group>

          <Form.Group className="mb-3">
            <Form.Label>Flight number</Form.Label>
            <Form.Control
              type="text"
              name="flightNumber"
              placeholder="flight nuber"
            />
          </Form.Group>

          <Form.Group className="mb-3">
            <Form.Label>Flight date</Form.Label>
            <Form.Control
              type="text"
              pattern="^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{2}$"
              placeholder="DD-MM-YY"
              name="date"
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Label>Aircraft type</Form.Label>
            <Form.Select aria-label="Default select example" name="aircraft">
              <option value="ATR">ATR</option>
              <option value="Airbus 320">Airbus 320</option>
              <option value="Boeing 737 Max">Boeing 737 Max</option>
            </Form.Select>
          </Form.Group>

          <Button variant="primary" type="submit">
            Generate Vouchers
          </Button>

        </Form>
      </Container>
    </>
  );
}

export default VoucherGeneratorForm;
