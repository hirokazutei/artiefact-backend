# -*- coding: utf-8 -*-
import sys
import messages
from main.models import session
from sqlalchemy.exc import IntegrityError
from main.models.factories import account

if __name__ == '__main__':
    with session.no_autoflush:
        user = account.create_user()
        try:
            session.commit()
        # Sometimes the names overlap
        except IntegrityError as error:
            print(error)
            print(messages.DUPLICATE_KEY_ERROR)
            sys.exit(0)
        else:
            account.UserAgreementFactory.create(user_id=user.id)
            account.ProfileFactory.create(user_id=user.id)
            try:
                session.commit()
            except Exception as error:
                print(error)
                sys.exit(0)
            else:
                print('created user: username={0}'.format(
                    user.username.username_lower,
                ))
