Hex arch + gRPC.

En este repo podemos encontrar un proyecto estructurado con arquitectura hexagonal y el protocolo gRPC.


**Arquitectura Hexagonal**

La Arquitectura Hexagonal, dada a conocer por Alistair Cockburn y también conocida como arquitectura de puertos y adaptadores, tiene como principal motivación separar nuestra aplicación en distintas capas o regiones con su propia responsabilidad. De esta manera consigue desacoplar capas de nuestra aplicación permitiendo que evolucionen de manera aislada. Además, tener el sistema separado por responsabilidades nos facilitará la reutilización.

Gracias a este desacoplamiento obtenemos también la ventaja de poder testear estas capas sin que intervengan otras externas, falseando el comportamiento con dobles de pruebas.


![image](https://github.com/LautaroOlmedo/Basic-Arithmetic/assets/72723456/652411b1-4fbc-4ace-a569-1354deb1cc5e)



**gRPC**

gRpc es un protocolo RPC (Remote Procedure Call), el cual puede ser ejecutado sobre cualquier entorno para realizar conexiones o llamadas entre microservicios. Este protocolo fue creado por Google y puede ser utilizado como una alternativa a REST, WebHook o GraphQL.
Entre sus principales características podemos encontrar:

* Transmisión bidireccional y autenticación conectable totalmente integrada con transporte basado en HTTP / 2.

* Destaca el rendimiento, debido a su bajo consumo de CPU, ancho de banda,etc.

* Ofrece JSON encondings y serialización con PROTO 3.

* Genera automáticamente stubs idiomáticos de cliente y servidor para su servicio en una variedad de idiomas y plataformas.

* Autenticación incorporada soportando SSL.

![image](https://github.com/LautaroOlmedo/Basic-Arithmetic/assets/72723456/bffbb04e-ca0b-4f23-b069-f2f921cfb4c3)







FUENTE:

Arquitectura Hexagonal: https://medium.com/@edusalguero/arquitectura-hexagonal-59834bb44b7f

gRPC: https://refactorizando.com/que-es-grpc/
