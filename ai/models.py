import tensorflow as tf
from tensorflow import keras
from keras import models
from keras.layers import Dense, Input
from keras.utils import to_categorical
from keras.datasets import mnist
from keras.utils.vis_utils import model_to_dot
import numpy as np
from simulation_pb2 import SimulationAction

from keras import backend as K

model_path='/home/param/repos/aimazedata/testv2_3/model.tf'
# simulate_ai - simple array features
def simulate_ai(features):
    numpyf = np.array(features)
    numpyf.shape=(1,10)
    in_data = tf.convert_to_tensor(numpyf, dtype=tf.float32)
    model = models.load_model(model_path)
    preds = model.predict(in_data)
    return preds

def get_action(sim, actionArr):
    maxidx = 0
    for i in range(len(actionArr)):
        if tf.math.greater(actionArr[i], actionArr[maxidx]):
            maxidx = i

    action = "up"

    if maxidx == 0:
        action="up"
    elif maxidx == 1:
        action="down"
    elif maxidx == 2:
        action="right"
    elif maxidx == 3:
        action="left"

    return SimulationAction(sim=sim, action=action)
