import pytest
import sys
sys.path.append('../')
from client import QuickNodeSDK


@pytest.fixture()
def client():
    return QuickNodeSDK()
