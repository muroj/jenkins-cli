/*
Author: Félix Belzunce Arcos
Since: January 2021
Description: Encode in Base64 all the credentials of a Jenkins Master at System level. The script is used to update all 
the credentials at System level, re-encrypting them in another Jenkins Master. 

The encoded message can be used to update the System credentials in a new Jenkins Master. 
*/

import com.cloudbees.plugins.credentials.SystemCredentialsProvider
import com.cloudbees.plugins.credentials.domains.DomainCredentials
import com.thoughtworks.xstream.converters.Converter
import com.thoughtworks.xstream.converters.MarshallingContext
import com.thoughtworks.xstream.converters.UnmarshallingContext
import com.thoughtworks.xstream.io.HierarchicalStreamReader
import com.thoughtworks.xstream.io.HierarchicalStreamWriter
import com.trilead.ssh2.crypto.Base64
import hudson.util.Secret
import hudson.util.XStream2
import jenkins.model.Jenkins

// Copy all domains from the system credentials provider
systemProvider = Jenkins.get().getExtensionList(SystemCredentialsProvider.class).first()
systemCreds = systemProvider.getDomainCredentials()
def stream = new XStream2()
stream.registerConverter(getDecrypter())
println stream.toXML(systemCreds).toString()

// The converter ensures that the output XML contains the unencrypted secrets
def getDecrypter() {
    SecretDecrypter = new Converter() {
        @Override
        boolean canConvert(Class type) { 
            type.equals(hudson.util.Secret.class) || type.equals(hudson.util.SecretBytes.class)
        }

        @Override
        void marshal(Object object, HierarchicalStreamWriter writer, MarshallingContext context) {
            writer.setValue(hudson.util.Secret.toString(object as Secret))
        }

        @Override
        Object unmarshal(HierarchicalStreamReader reader, UnmarshallingContext context) { null }
    }
}