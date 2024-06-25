import {features} from "../constants";
const Features = () => {
  return (
    <div className="relative mt-20 border-b border-neutral-800 min-h-[800px]">
        <div className="text-center">
            <span className="bg-neutral-900 text-sky-500 rounded-full h-6 text-lg font-medium px-2 py-1 uppercase">
                features
            </span>
            <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide">
                Test the prototype of blockchain-based 
                <span className="flex justify-center bg-gradient-to-r from-sky-800 to-sky-500 text-transparent bg-clip-text">
                    identity manage application
                </span> 
            </h2>
        </div>
        <div className="flex flex-wrap mt-10 lg:mt-20">
            {features.map((feature,index) => (

                <div key={index} className="w-full sm:1/2 lg:w-1/3">
                    <div className="flex">
                        <div className="flex mx-6 h-10 w-10 p-2 bg-neutral-900 text-sky-500 justify-center items-center rounded-full">
                            {feature.icon}
                        </div>
                        <div>
                            <h5 className="mt-1 mb-6 text-xl">
                                {feature.text}
                            </h5>
                            <p className="text-md p-2 mb-20 text-neutral-500">
                               {feature.description} 
                            </p>
                        </div>
                    </div>
                </div>

            ))}
        </div>
    </div> 
  )
}

export default Features