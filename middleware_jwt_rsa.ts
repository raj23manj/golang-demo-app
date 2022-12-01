import { Request, Response, NextFunction } from 'express';
import { Configuration, ConfigOptions } from './../config/configuration.config';
import jwt from 'jsonwebtoken';
const jwksClient = require('jwks-rsa');
// add a property to request object for typescript to identify
declare global {
  namespace Express {
    interface Request {
      currentUser?: any | null;
    }
  }
}

export const currentUser = async (
  req: Request,
  _res: Response,
  next: NextFunction,
) => {
  // if not authorization header found return setting a null
  // in your project handle the case how ever you want adding a
  // middleware
  if(!req.headers?.authorization) {
    req.currentUser = null;
    return next();
  }
  // verification of the jwt kc token code
  const token = (req.headers?.authorization! as string).replace('Bearer ', '');
  // get the tenant name from the iss
  const tenant = JSON.parse(atob(token.split('.')[1])).iss.split('/').slice(-1)[0];
  const config: ConfigOptions = Configuration.getConfig();
  const client = jwksClient({
    jwksUri: `${config.domain}/v2/auth/${tenant}/certs`
  });
  const options: any = {
    issuer: `${config.domain}/auth/realms/${tenant}`,
    algorithms: ['RS256']
  };

  function getKey(header: any, callback: any) {
    client.getSigningKey(header.kid, function(err: any, key: { publicKey: any; rsaPublicKey: any; }) {
      var signingKey = key.publicKey || key.rsaPublicKey;
      callback(null, signingKey);
    });
  }

  jwt.verify(token, getKey, options, async(err, decoded: any) => {
    if(err) {
      console.log("verification:error:", err);
      req.currentUser = null;
    } else {
      // add additional parameters if needed from decoded
      const { type, fp_identifier, tenant} = decoded
      req.currentUser = {
        type: type,
        fpIdentifier: fp_identifier,
        tenant: tenant
      };
    }
    next();
  });
};
