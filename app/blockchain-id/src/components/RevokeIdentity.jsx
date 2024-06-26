import { useState } from "react";
import perfil from "../assets/perfil.png"
import QRCode from 'qrcode.react';
import {
  Card,
  Input,
  Checkbox,
  Button,
  Typography,
  CardBody,
  CardFooter,
} from "@material-tailwind/react";

const RevokeIdentity = () => {
  return (
    <div className="mt-20 border-t border-neutral-800">
        <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide text-center">
            Test the Revoke Indentity  
            <span className="flex justify-center bg-gradient-to-r from-sky-800 to-sky-500 text-transparent bg-clip-text">
                Smart Contract
            </span> 
        </h2>
        <div className="flex flex-wrap items-center justify-center mt-10">
            <div className="p-2 w-full lg:w-1/2">
      <Card color="transparent" shadow={false} className="w-full flex items-center">
      <Typography color="gray" className="mt-1 font-normal">
        Nice to meet you! Enter the id to revoke identity.
      </Typography>
      <form  className="mt-8 mb-2 w-80 max-w-screen-lg sm:w-96">
        <div className="mb-1 flex flex-col gap-6">
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            ID to revoke
          </Typography>
          <Input
            type= "text"
            size="lg"
            placeholder="ex: 201820983"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
        </div>
        <Button type="submit" className="mt-6 bg-gradient-to-r from-sky-500 to-sky-800 transform transition duration-500 hover:scale-125" fullWidth>
          Revoke Identity
        </Button>
      </form>
    </Card>
    </div>
    <div className="p-2 w-full lg:w-1/2">

    </div>
        </div>
    </div>
  )
}

export default RevokeIdentity