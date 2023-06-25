Hex arch + gRPC.

En este repo podemos encontrar un proyecto estructurado con arquitectura hexagonal y el protocolo gRPC


**Arquitectura Hexagonal**

La Arquitectura Hexagonal, dada a conocer por Alistair Cockburn y también conocida como arquitectura de puertos y adaptadores, tiene como principal motivación separar nuestra aplicación en distintas capas o regiones con su propia responsabilidad. De esta manera consigue desacoplar capas de nuestra aplicación permitiendo que evolucionen de manera aislada. Además, tener el sistema separado por responsabilidades nos facilitará la reutilización.

Gracias a este desacoplamiento obtenemos también la ventaja de poder testear estas capas sin que intervengan otras externas, falseando el comportamiento con dobles de pruebas.


![hexagonal](https://github.com/LautaroOlmedo/Basic-Arithmetic/assets/72723456/e43fe007-20a4-4604-8390-492640798ed5)


**gRPC**

gRpc es un protocolo RPC (Remote Procedure Call), el cual puede ser ejecutado sobre cualquier entorno para realizar conexiones o llamadas entre microservicios. Este protocolo fue creado por Google y puede ser utilizado como una alternativa a REST, WebHook o GraphQL.
Entre sus principales características podemos encontrar:

* Transmisión bidireccional y autenticación conectable totalmente integrada con transporte basado en HTTP / 2.

* Destaca el rendimiento, debido a su bajo consumo de CPU, ancho de banda,etc.

* Ofrece JSON encondings y serialización con PROTO 3.

* Genera automáticamente stubs idiomáticos de cliente y servidor para su servicio en una variedad de idiomas y plataformas

* Autenticación incorporada soportando SSL.

![image](https://github.com/LautaroOlmedo/Basic-Arithmetic/assets/72723456/16b1d73d-3776-452a-8dd4-b779ce93e247)






FUENTE:

Arquitectura Hexagonal: https://medium.com/@edusalguero/arquitectura-hexagonal-59834bb44b7f

gRPC: https://refactorizando.com/que-es-grpc/
