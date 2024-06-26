import { useState } from "react";
import perfil from "../assets/perfil.png"
import {
  Card,
  Input,
  Checkbox,
  Button,
  Typography,
  CardBody,
  CardFooter,
} from "@material-tailwind/react";

const IssueIdentity = () => {

        const [formData, setFormData] = useState({
          nombre: '',
          apellido: '',
          fechaNacimiento: '',
          numeroCarne: '',
          direccion: '',
          telefono: '',
          email: '',
        });
      
        const handleChange = (e) => {
          setFormData({ ...formData, [e.target.name]: e.target.value });
        };
      
        const handleSubmit = async (e) => {
          e.preventDefault();
          try {
            const response = await fetch('URL_DE_TU_API', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify(formData),
            });
            if (response.ok) {
              alert('Registro exitoso');

            } else {
              alert('Error al registrar');
            }
          } catch (error) {
            console.error('Error al registrar:', error);
            alert('Error al registrar');
          }
        };

  return (
    <div className="mt-20">
        <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide text-center">
            Test the Indentity Issue  
            <span className="flex justify-center bg-gradient-to-r from-sky-800 to-sky-500 text-transparent bg-clip-text">
                Smart Contract
            </span> 
        </h2>
        <div className="flex flex-wrap items-center justify-center mt-10">
            <div className="p-2 w-full lg:w-1/2">
      <Card color="transparent" shadow={false} className="w-full flex items-center">
      <Typography color="gray" className="mt-1 font-normal">
        Nice to meet you! Enter your details to issue your new identity.
      </Typography>
      <form onSubmit={handleSubmit} className="mt-8 mb-2 w-80 max-w-screen-lg sm:w-96">
        <div className="mb-1 flex flex-col gap-6">
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Your Name
          </Typography>
          <Input
            size="lg"
            placeholder="ex: Juan Carlos"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Your Last Name
          </Typography>
          <Input
            size="lg"
            placeholder="ex: Doe Ramirez"
            className=" !border-t-blue-sky-500 focus:!border-t-sky-800"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            University
          </Typography>
          <Input
            size="lg"
            placeholder="ex: EPN"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Your Faculty
          </Typography>
          <Input
            size="lg"
            placeholder="ex: Sistemas"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Semester
          </Typography>
          <Input
            size="lg"
            placeholder="ex: First"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Blood Type
          </Typography>
          <Input
            size="lg"
            placeholder="ex: A+"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Address
          </Typography>
          <Input
            size="lg"
            placeholder="ex: Av. Ladrón de Guevara"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
        </div>
        <Button type="submit" className="mt-6 bg-gradient-to-r from-sky-500 to-sky-800" fullWidth>
          Save Data
        </Button>
      </form>
    </Card>
    </div>
    <div className="p-2 w-full lg:w-1/2 mt-0 flex justify-center text-center">
    <Card className="shadow-lg shadow-sky-500/100 rounded-lg w-96 p-5">
      <CardBody>
      <div className="w-full flex justify-center">
        <img src={perfil} alt="default-profile-picture" className="w-44"/>
      </div>
        <Typography className="mb-2 mt-10 text-2xl">
          Data on Blockchain
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
      </CardBody>
      <CardFooter className="pt-0">
      </CardFooter>
    </Card>
    </div>
        </div>
    </div>
  )
}

export default IssueIdentity