написать 2 микросервиса которые взаимодействуют по grpc.
на одном разместить простую веб страницу с кнопкой и текстовым полем.
на втором разместить бизнес логику игры.
при нажатии на кнопку сервис 1 делает запрос на игру в сервис 2, получает ответ и показывает его в текстовом поле.
Результат форматируем всегда с 2 десятичными знаками, для 0 это 0.00 .
логика для игры очень простая, на каждый запрос на игру мы:
генерируем случайное число от 0 до 1 с 2мя десятичными знаками (например 0.22, 0.35 и тд).
результатом игры считается сгенерированное значение сложенное с предыдущим результатом игры (результатом с предыдущего
запроса). Для первой игры предыдущий результат равен 0.
мы возвращаем результат в сервис 1.

требования к результату тестового задания
желательно куда-то задеплоить и скинуть ссылку на готовое решение
код разместить на любой VCS какую используете или нравится
описать в пояснительной записке какие бы вы предложили решения из области system design для того чтобы заскейлить эту
игру на скажем 10000 rps (решать не надо! 🙂

требования к коду
идеоматичность - соответствование общепринятым идиомам разработки на языке Go
использование какой либо подхода к архитектуре кодовой базы ( ddd, clean code, 3 layers, mvс - все что угодно но описать
это в записке к тестовому заданию или в комментариях).
написать хотя бы один интеграционный тест. юнит тесты можно опустить.
предусмотреть обработку ошибок( ошибку можно не выводить ).
хранение результата игры - все что угодно (вплоть до текстового файла или in memory).