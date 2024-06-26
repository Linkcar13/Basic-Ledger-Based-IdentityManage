import { useState } from "react";
import perfil from "../assets/usuario.png"
import {
  Card,
  Input,
  Checkbox,
  Button,
  Typography,
  CardBody,
  CardFooter,
} from "@material-tailwind/react";

const ValidateIdentity = () => {
  return (
    <div className="mt-20 border-t border-neutral-800">
        <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide text-center">
            Test the Indentity Validate  
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
      <form  className="mt-8 mb-2 w-80 max-w-screen-lg sm:w-96">
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
            Your Faculty
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
          Verify Identity
        </Button>
      </form>
    </Card>
    </div>
    <div className="p-2 w-full lg:w-1/2">
    <Card className="mt-0 w-full flex items-center">
      <CardBody>
        <img src={perfil} alt="default-profile-picture" className="w-48"/>
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

export default ValidateIdentity
