import { StatusCodes } from "http-status-codes"

export const errorHandlerMiddleware = (err,req,res,next) =>{
    console.log(err)
    res.status(500).json({msg:err})
}

    //const defaultError = {
    //    statusCode: StatusCodes.INTERNAL_SERVER_ERROR,
    //   msg:'Something went wrong, try again later',
    //
    //if(err.name == 'ValidationError'){
    //    defaultError.statusCode = StatusCodes.BAD_REQUEST
    //    defaultError.msg = err.message
    
   //res.status(defaultError.statusCode).json({msg: err})
   //res.status(defaultError.statusCode).json({msg: defaultError.msg})
