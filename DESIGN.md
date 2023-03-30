# Camera review design
## Overview
<p>Una empresa requiere un sistema que permita a fotógrafos profesionales el subir reviews de 
cámaras fotográficas para que cualquier persona desde cualquier parte del mundo pueda buscar los reviews y comprarlas a través de su portal. La empresa cuenta con un equipo de developers especializado en frontend que realizará un portal para que los editores suban los reviews y los usuarios puedan verlos, y ha solicitado que el especialista en backend les proporcione </p>

## Alcance
### Casos de uso
- Como editor quiero subir una review de una cámara.
- Como editor quiero subir una review de una lente.
- Como usuario no registrado quiero poder leer una review.
### Casos de uso fuera del scope
- Como usuario no registrado quiero subir una review de cámara.
### Modelado de datos
Review:
- title
- content
- author
- product
Product
- name
- brand
- description
User
- name
- email
- lastname
- createdAt
- isAdmin
- isStaff
